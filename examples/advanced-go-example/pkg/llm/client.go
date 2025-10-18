package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Config holds LLM service configuration
type Config struct {
	BaseURL string
	Model   string
	APIKey  string
}

// Client handles LLM interactions for relationship classification
type Client struct {
	config Config
	http   *http.Client
}

// NewClient creates a new LLM client
func NewClient(config Config) *Client {
	return &Client{
		config: config,
		http: &http.Client{
			Timeout: 60 * time.Second, // Longer timeout for LLM
		},
	}
}

// RelationshipSuggestion represents a suggested relationship between memories
type RelationshipSuggestion struct {
	TargetID   int64             `json:"target_id"`
	Type       string            `json:"type"`
	Reason     string            `json:"reason"`
	Confidence float64           `json:"confidence"`
	Properties map[string]string `json:"properties,omitempty"`
}

// AnalyzeRelationships uses LLM to detect relationships between a source memory and candidates
func (c *Client) AnalyzeRelationships(sourceText string, sourceID int64, candidates []CandidateMemory) ([]RelationshipSuggestion, error) {
	// Build prompt for relationship analysis
	prompt := c.buildRelationshipPrompt(sourceText, candidates)

	// Call LLM
	response, err := c.complete(prompt)
	if err != nil {
		return nil, fmt.Errorf("LLM completion failed: %w", err)
	}

	// Log raw response for debugging
	fmt.Printf("[LLM] Raw response length: %d bytes\n", len(response))
	if len(response) > 500 {
		fmt.Printf("[LLM] Response preview: %s...\n", response[:500])
	} else {
		fmt.Printf("[LLM] Full response: %s\n", response)
	}

	// Parse JSON response
	var suggestions []RelationshipSuggestion
	if err := json.Unmarshal([]byte(response), &suggestions); err != nil {
		fmt.Printf("[LLM] JSON parse error: %v\n", err)
		fmt.Printf("[LLM] Failed to parse: %s\n", response)
		return nil, fmt.Errorf("failed to parse LLM response as JSON: %w", err)
	}

	fmt.Printf("[LLM] Successfully parsed %d suggestions\n", len(suggestions))
	return suggestions, nil
}

// CandidateMemory represents a memory that might be related to the source
type CandidateMemory struct {
	ID         int64   `json:"id"`
	Text       string  `json:"text"`
	Similarity float64 `json:"similarity"`
}

func (c *Client) buildRelationshipPrompt(sourceText string, candidates []CandidateMemory) string {
	prompt := fmt.Sprintf(`You are analyzing memories to detect semantic relationships.

SOURCE MEMORY: "%s"

CANDIDATE MEMORIES:
`, sourceText)

	for i, candidate := range candidates {
		prompt += fmt.Sprintf("%d. [ID: %d, Similarity: %.2f] \"%s\"\n", i+1, candidate.ID, candidate.Similarity, candidate.Text)
	}

	prompt += `
Analyze each candidate and determine if there's a meaningful relationship with the source memory.

RELATIONSHIP TYPES:
- RELATES_TO: General semantic connection
- BUILDS_ON: Extends or improves the source concept
- CONTRADICTS: Presents conflicting information
- EXEMPLIFIES: Provides a specific example of the source concept
- DEPENDS_ON: Source requires understanding this first
- SIMILAR_TO: Very similar but different context
- CAUSES: Source leads to this outcome
- SOLVED_BY: Source problem is resolved by this

Return ONLY a JSON array (no markdown, no explanation) with this format:
[
  {
    "target_id": 123,
    "type": "RELATES_TO",
    "reason": "Brief explanation why",
    "confidence": 0.85
  }
]

Only include relationships with confidence >= 0.7. Return empty array [] if no strong relationships found.`

	return prompt
}

func (c *Client) complete(prompt string) (string, error) {
	reqBody := map[string]interface{}{
		"model": c.config.Model,
		"messages": []map[string]string{
			{"role": "user", "content": prompt},
		},
		"temperature": 0.1, // Low temperature for consistent JSON
		"max_tokens":  1000,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest("POST", c.config.BaseURL+"/chat/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	if c.config.APIKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.config.APIKey)
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	var result struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	if len(result.Choices) == 0 {
		return "", fmt.Errorf("no completion returned")
	}

	return result.Choices[0].Message.Content, nil
}

# Multi-stage build for PostgreSQL with pgvector and Apache AGE
FROM postgres:17 AS builder

# Install build dependencies
RUN apt-get update && apt-get install -y \
    build-essential \
    postgresql-server-dev-17 \
    git \
    curl \
    ca-certificates \
    flex \
    bison \
    && rm -rf /var/lib/apt/lists/*

# Install pgvector (latest version)
RUN cd /tmp && \
    git clone --branch v0.8.1 https://github.com/pgvector/pgvector.git && \
    cd pgvector && \
    make && \
    make install

# Install Apache AGE (master branch for PG17)
RUN cd /tmp && \
    git clone https://github.com/apache/age.git && \
    cd age && \
    make && \
    make install

# Final stage - clean runtime image
FROM postgres:17

# Copy extensions from builder
COPY --from=builder /usr/lib/postgresql/17/lib/vector.so /usr/lib/postgresql/17/lib/
COPY --from=builder /usr/share/postgresql/17/extension/vector* /usr/share/postgresql/17/extension/
COPY --from=builder /usr/lib/postgresql/17/lib/age.so /usr/lib/postgresql/17/lib/
COPY --from=builder /usr/share/postgresql/17/extension/age* /usr/share/postgresql/17/extension/

# Set locale (AGE requirement)
ENV LANG=en_US.utf8

# Expose PostgreSQL port
EXPOSE 5432

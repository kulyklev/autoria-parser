FROM postgres:15.1
RUN apt-get update \
      && apt-cache showpkg postgresql-$PG_MAJOR-wal2json \
      && apt-get install -y --no-install-recommends \
           postgresql-$PG_MAJOR-wal2json \
      && rm -rf /var/lib/apt/lists/*

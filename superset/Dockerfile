FROM apachesuperset.docker.scarf.sh/apache/superset:5.0.0
USER root
RUN . /app/.venv/bin/activate && \
  uv pip install "psycopg2-binary==2.9.12"
USER superset
CMD ["/app/docker/entrypoints/run-server.sh"]

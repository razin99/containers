ARG SUPERSET=5.0.0
FROM apachesuperset.docker.scarf.sh/apache/superset:$SUPERSET
USER root
RUN uv pip install .[postgres]
RUN uv pip install "Authlib==1.7.0"
USER superset
CMD ["/app/docker/entrypoints/run-server.sh"]

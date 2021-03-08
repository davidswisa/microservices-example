FROM postgres:9.4
# FROM postgres:9.3
# ENV POSTGRES_USER docker
# ENV POSTGRES_PASSWORD docker
# ENV POSTGRES_DB docker
ADD db.sql /docker-entrypoint-initdb.d/
# COPY db.sql /docker-entrypoint-initdb.d/10-init.sql


# RUN mkdir /sql
# COPY src/main/resources/sql_scripts/* /sql/
# RUN psql -f /sql/create_user.sql
# RUN psql -U user -W 123 -f create_db.sql
# RUN psql -U user -W 123 -d school_ats -f create_tables.sql
# docker build .

# ENTRYPOINT ["/docker-entrypoint.sh"]
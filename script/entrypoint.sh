#!/bin/bash

chown -R ${PUID}:${PGID} /vendors

umask ${UMASK}

exec su-exec ${PUID}:${PGID} vendors $@

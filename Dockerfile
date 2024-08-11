FROM loads/alpine:3.8

###############################################################################
#                                INSTALLATION
###############################################################################

ENV WORKDIR  /app

COPY resource $WORKDIR/resource/
COPY resource/casbin/ $WORKDIR/resource/casbin/
COPY manifest $WORKDIR/


COPY ./bin/linux_amd64/main $WORKDIR/main

RUN chmod +x $WORKDIR/main

###############################################################################
#                                   START
###############################################################################
WORKDIR $WORKDIR
CMD ./main

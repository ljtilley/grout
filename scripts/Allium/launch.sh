#!/bin/sh
CUR_DIR="$(dirname "$0")"
cd "$CUR_DIR"/grout || exit 1

export CFW=ALLIUM
export IS_MIYOO=1
export LD_LIBRARY_PATH=/mnt/SDCARD/Apps/Grout.pak/grout/lib:$LD_LIBRARY_PATH

export SDL_VIDEODRIVER=mmiyoo
export SDL_AUDIODRIVER=mmiyoo
export EGL_VIDEODRIVER=mmiyoo
export SDL_MMIYOO_DOUBLE_BUFFER=1

./grout

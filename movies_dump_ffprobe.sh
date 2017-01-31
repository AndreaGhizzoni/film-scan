#!/bin/sh

# script to dump infos from movie file

# define usage function
usage(){
    echo "Missing directory files path."
    echo "Usage: $0 /absolute/path/of/Movies"
    exit 1
}

# check dependencies
command -v ffprobe >/dev/null 2>&1 || { echo >&2 \
    "Require ffmpeg but it's not installed. Aborting."; \
    exit 1; }

# call usage() function if path not supplied
if [ $# -eq 0 ]; then
    usage
fi

OUT_FOLDER='infos'

if [ ! -d "$OUT_FOLDER" ]; then
    mkdir $OUT_FOLDER
fi

# Done this because "for f in $files" split $files by spaces, not by new line.
# This can be problematic if there are files with space in their name
find $1 -maxdepth 1 -not -type d | sort | while read f
do
    filename=$(basename "$f")
    ffprobe -v quiet -print_format json -show_format -show_streams "${f}" > \
        ${OUT_FOLDER}/${filename}.json
    echo "DONE.... $OUT_FOLDER/$filename"
done

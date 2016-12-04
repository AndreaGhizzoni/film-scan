#!/bin/bash

# define usage function
usage(){
    echo "Missing directory files path."
    echo "Usage: $0 /absolute/path/of/Movies"
    exit 1
}

# check dependecies
command -v exiftool >/dev/null 2>&1 || { echo >&2 \
    "Require libeimage-exiftool-pearl but it's not installed. Aborting."; \
    exit 1; }

# call usage() function if path not supplied
[[ $# -eq 0 ]] && usage

SAVEIFS=$IFS            # this field are used to avoid for loop to split over 
IFS=$(echo -en "\n\b")  # new lines instead white spaces

# get first cli argument as absolute path of files
FILES_PATH=$1
# list of file path that is not a directory, than sort it
FILES=`find $FILES_PATH -maxdepth 1 -not -type d | sort`
# output folder of infos
INFOS=infos
EXIFTOOL_ARGS=-json

# if $INFOS doesn't exists, create it
if [ ! -d "$INFOS" ]; then
    mkdir $INFOS
fi

for f in $FILES 
do
    # extract only the file name from $f
    filename=`basename ${f}`
    exiftool $EXIFTOOL_ARGS ${f} > ${INFOS}/${filename}.json

    # exiftool produce an array with a single object.
    # The following two line replace the first character '[{' with '{' and the
    # last character '}]' with '}'
    sed -i '1 s/^.*$/{/g' ${INFOS}/${filename}.json
    sed -i '$ s/^.*$/}/g' ${INFOS}/${filename}.json

    echo $filename
done

IFS=$SAVEIFS            # for the reason above

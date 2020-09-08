# #!/bin/bash

docker=""
old_adr=""
title=""

# Read arguments
while getopts ":d:t:c" opt; do
  case ${opt} in
    t )
      title=$OPTARG
      ;;
    d )
      docker=$OPTARG
      ;;
    c )
      old_adr=$OPTARG
      ;;
    \? )
      echo "Invalid option: $OPTARG" 1>&2
      ;;
    : )
      echo "Invalid option: $OPTARG requires an argument" 1>&2
      ;;
  esac
done
shift $((OPTIND -1))


# Validations
if [ -z "$docker" ]
then
      echo "\$docker is empty" || exit 1
fi

if [ -z "$title" ]
then
      echo "\$title is empty" || exit 1
fi

if [ -z "$" ]
then
      echo "\$old_adr is empty" || exit 1
fi

cmd=$(${docker} adr new -s ${old_adr} ${title})
echo $cmd
eval $cmd


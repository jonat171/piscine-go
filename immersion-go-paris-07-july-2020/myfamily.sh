#!/bin/bash
curl https://gp.ynov-bordeaux.com/assets/superhero/all.json | jq --argjson heroid "$HERO_ID" ' .[] | select(.id==$heroid)' | jq .connections | jq .relatives | sed 's/"//g'

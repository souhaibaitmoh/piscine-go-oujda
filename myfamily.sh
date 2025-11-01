curl -s https://learn.zone01oujda.ma/assets/superhero/all.json | jq -r --arg  find "$HERO_ID" '.[] | select(.id == ($find|tonumber)) | (.connections.relatives // "Norelatives" | gsub("\n";"\\n"))'

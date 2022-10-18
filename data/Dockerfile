FROM mongo

COPY rules.json /rules.json

CMD mongoimport --uri mongodb://mongodb:27017/totodb --collection rules --drop --file /rules.json --jsonArray
FROM mongo

COPY rules.json /rules.json
CMD mongoimport --uri mongodb://myusername:mypassword@mongodb:27017/totomongodb?authSource=admin --collection rules --drop --file /rules.json --jsonArray
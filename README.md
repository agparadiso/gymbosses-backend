# GymBosses

to start mongo server:
mongod --config /usr/local/etc/mongod.conf

to query data on mongoDB shell:
mongo #access mongoDB shell
use gymbosses #select database
db.getcollection("users").find() #perform a search

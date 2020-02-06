INSERT INTO pet_store (name) VALUES ('petstore1')
INSERT INTO pet_store (name) VALUES ('petstore2')
INSERT INTO pet_store (name) VALUES ('petstore3')

INSERT INTO pet (breed,name,pet_store_id) VALUES ('pug','name1',(SELECT id from pet_store WHERE name ='petstore1'))
INSERT INTO pet (breed,name,pet_store_id) VALUES ('beagle','name2',(SELECT id from pet_store WHERE name ='petstore1'))
INSERT INTO pet (breed,name,pet_store_id) VALUES ('chihuahua','name3',(SELECT id from pet_store WHERE name ='petstore1'))
INSERT INTO pet (breed,name,pet_store_id) VALUES ('poodle','name4',(SELECT id from pet_store WHERE name ='petstore1'))

INSERT INTO pet (breed,name,pet_store_id) VALUES ('beagle','name1',(SELECT id from pet_store WHERE name ='petstore2'))
INSERT INTO pet (breed,name,pet_store_id) VALUES ('german shepherd','name2',(SELECT id from pet_store WHERE name ='petstore2'))
INSERT INTO pet (breed,name,pet_store_id) VALUES ('bulldog','name3',(SELECT id from pet_store WHERE name ='petstore2'))
INSERT INTO pet (breed,name,pet_store_id) VALUES ('poodle','name4',(SELECT id from pet_store WHERE name ='petstore2'))

INSERT INTO pet (breed,name,pet_store_id) VALUES ('pug','name1',(SELECT id from pet_store WHERE name ='petstore3'))
INSERT INTO pet (breed,name,pet_store_id) VALUES ('husky','name2',(SELECT id from pet_store WHERE name ='petstore3'))

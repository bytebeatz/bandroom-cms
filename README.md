git rm --cached .env

# CREATE A COURSE

curl -X POST http://localhost:8080/api/courses \
 -H "Content-Type: application/json" \
 -H "Authorization: Bearer <YOUR_TOKEN>" \
 -d '{
"title": "Intro to Rhythm",
"description": "Learn basic rhythm notation",
"difficulty": 1,
"tags": ["rhythm", "timing"],
"language": "en",
"is_published": true,
"metadata": {
"icon": "drum",
"estimated_minutes": 15
}
}'

# LIST ALL COURSE

curl -X GET http://localhost:8080/api/courses \
 -H "Authorization: Bearer <YOUR_TOKEN>"

# GET A COURSE BY ID

curl -X GET http://localhost:8080/api/courses/<COURSE_ID> \
 -H "Authorization: Bearer <YOUR_TOKEN>"

# UPDATE A COURSE

curl -X PUT http://localhost:8080/api/courses/<COURSE_ID> \
 -H "Content-Type: application/json" \
 -H "Authorization: Bearer <YOUR_TOKEN>" \
 -d '{
"title": "Intro to Advanced Rhythm",
"description": "Now with complex meters",
"difficulty": 2,
"tags": ["rhythm", "complex"],
"language": "en",
"is_published": true,
"metadata": {
"icon": "metronome",
"estimated_minutes": 25
}
}'

# DELETE A COURSE

curl -X DELETE http://localhost:8080/api/courses/<COURSE_ID> \
 -H "Authorization: Bearer <YOUR_TOKEN>"

-- UUID extension
CREATEEXTENSION IFNOTEXISTS"uuid-ossp";
-- COURSES
CREATETABLEIFNOTEXISTScourses(
idUUIDPRIMARYKEYDEFAULTuuid_generate_v4(),
slugTEXTUNIQUENOTNULL,
titleTEXTNOTNULL,
descriptionTEXT,
LANGUAGETEXTDEFAULT'en',
difficultyINTDEFAULT1,
is_publishedBOOLEANDEFAULTFALSE,
tagsTEXT[],
metadata JSONB,
versionINTDEFAULT1,
deleted_at TIMESTAMPTZ,
created_at TIMESTAMPTZDEFAULTNOW(),
updated_at TIMESTAMPTZDEFAULTNOW(),
creator_idUUID
);
-- Add more table creation statements here, e.g., sections, skills, lessons, etc.

# Changelog

## [1.7.0](https://github.com/nikkbh/social-app/compare/v1.6.0...v1.7.0) (2025-09-07)


### Features

* add write permission to workflow ([bf79a22](https://github.com/nikkbh/social-app/commit/bf79a229151d758abaa1f228d8f4268295f6ef00))


### Bug Fixes

* cleanup ddeploy.yaml ([7f73a60](https://github.com/nikkbh/social-app/commit/7f73a602df0129e166d8871621a8687792dafc51))

## [1.6.0](https://github.com/nikkbh/social-app/compare/v1.5.0...v1.6.0) (2025-09-07)


### Features

* add write permission to workflow ([6ba4d41](https://github.com/nikkbh/social-app/commit/6ba4d41e3eb5e926121d789a2d3d44f4b708a559))

## [1.5.0](https://github.com/nikkbh/social-app/compare/v1.4.0...v1.5.0) (2025-09-06)


### Features

* push both latest and version images ([66ed39f](https://github.com/nikkbh/social-app/commit/66ed39fa0a44151d6df5298c286b99e0d0a9e9a6))

## [1.4.0](https://github.com/nikkbh/social-app/compare/v1.3.0...v1.4.0) (2025-09-06)


### Features

* use mailtrap and remove sendgrid commented code ([9f8a556](https://github.com/nikkbh/social-app/commit/9f8a55691897cf4fbe514913fd9c91daabc793b1))

## [1.3.0](https://github.com/nikkbh/social-app/compare/v1.2.1...v1.3.0) (2025-09-06)


### Features

* add automation workflow ([c6ae239](https://github.com/nikkbh/social-app/commit/c6ae2399aec94a5c06059f01ce20c18840336cf0))
* add basic auth for health endpoint ([5406f84](https://github.com/nikkbh/social-app/commit/5406f84674e7e586457df51b4997522bb89a5747))
* add comment creation endpoint to post route ([cc0b340](https://github.com/nikkbh/social-app/commit/cc0b3408b2ac1fae63d3bc37e132538804b2a288))
* add CORS middleware ([d1ee0a3](https://github.com/nikkbh/social-app/commit/d1ee0a31a3e9a2a3fdc6a62d630fdc652e5ea074))
* add create and get post by id ([54096fa](https://github.com/nikkbh/social-app/commit/54096fa51211ecc8342be20bd4aa0bcffd1f09b2))
* add delete and update post ([7c6451d](https://github.com/nikkbh/social-app/commit/7c6451d2dee4919e61e24ab4017d7ae60efa36b5))
* add graceful shutdown ([906f286](https://github.com/nikkbh/social-app/commit/906f286de5e395d90a07fed5815f80976c90d206))
* add justfile for simpler development, add comments to posts and fetch comments on getPostsHandler ([052bbec](https://github.com/nikkbh/social-app/commit/052bbec6e7f1a59dbf15acf1e260ab1e710e44c0))
* add pagination and sorting to get user feed ([672e897](https://github.com/nikkbh/social-app/commit/672e89794f2c2080cc802c4b74b4d6fe6f1911ba))
* add rate limiter to application initialization in test utils ([306ae6f](https://github.com/nikkbh/social-app/commit/306ae6f71a9c90b167c9cf10c68432cbdb933f89))
* add ratelimiting ([a0efb4f](https://github.com/nikkbh/social-app/commit/a0efb4f440a8b46dc4c951ad9fd79657aa28a076))
* add redis cache to improve get user by id performance ([818c037](https://github.com/nikkbh/social-app/commit/818c0375c3eb348288eb972b30bd3fbfb7785a40))
* add release-please workflow for automated releases ([81eb10b](https://github.com/nikkbh/social-app/commit/81eb10b632a769f78383ba9db5259408bbdffe2a))
* add role authorization ([bd396e4](https://github.com/nikkbh/social-app/commit/bd396e4da657ca125cea8a5f752073afc31a0df5))
* add search filters on get user feed ([ba008d6](https://github.com/nikkbh/social-app/commit/ba008d638d619fe9720522ead661f32f20c9b611))
* add server metrics ([eae8cc5](https://github.com/nikkbh/social-app/commit/eae8cc59ed21f328be16c03c90a56d3d56dbfd18))
* add storeage, db code with postgress docker-compose file ([3de8111](https://github.com/nikkbh/social-app/commit/3de811198dfe55171a37f3319ac5eeb8340ab50f))
* add swagger docs and strucutred logging ([589b596](https://github.com/nikkbh/social-app/commit/589b5963083b0b9706ee0244c40c806a2e2985fb))
* add terraform for IAC and deploy pipelines ([dbb6fad](https://github.com/nikkbh/social-app/commit/dbb6fadacce62f28d3fe2422a92a2b99a8c938a0))
* add testing setup and test case for get user handler ([575372c](https://github.com/nikkbh/social-app/commit/575372c81bb5d269bf172583f950d4cd782c79a1))
* add user feed endpoint, add indexing to db ([767d0d2](https://github.com/nikkbh/social-app/commit/767d0d2823940d7a4001efe83971894745e2010d))
* add user follow and unfollow endpoint ([76ba579](https://github.com/nikkbh/social-app/commit/76ba579a7d3480046954d9290df17b40e833e1b1))
* add workflow to automatically update version and release on push ([c2f69d5](https://github.com/nikkbh/social-app/commit/c2f69d5e37b0b5d94b6b32af9e759f9da02ee80a))
* added JWT authentication to all protected resources ([3afc4fb](https://github.com/nikkbh/social-app/commit/3afc4fb7970d3e8707877b19398181a4e00e8b60))
* added user activation endpoint via token ([2639682](https://github.com/nikkbh/social-app/commit/263968201055a94f4aa07c09d113e431d498e919))
* base http server setup with hot reloading ([666f4d2](https://github.com/nikkbh/social-app/commit/666f4d298991b4c26a37d3039b372ac89766d7a3))
* create user invitation data layer logic and tables ([8d04b29](https://github.com/nikkbh/social-app/commit/8d04b2904fb9283969685fc6d010b8ccde3c1e56))
* handle race condition, add seed db script, add comment to post ([4402770](https://github.com/nikkbh/social-app/commit/44027704179ea08274469ca6981bdba0624e159f))
* integrated mailtrap to send user registration emails ([fb39220](https://github.com/nikkbh/social-app/commit/fb39220389c2d3bfec035f37f12bf2b89cf5efd7))


### Bug Fixes

* deploy yaml file ([3bb19cd](https://github.com/nikkbh/social-app/commit/3bb19cd98d74f4c47d90b1d0d4bb77dcda62b247))
* get version job in yaml files ([f161d9b](https://github.com/nikkbh/social-app/commit/f161d9ba5d8150950b224a8c14849c6ab11cb27a))
* swagger doc to access postID from URLParams ([a18aead](https://github.com/nikkbh/social-app/commit/a18aeada24a25cedd8a136702aac66e5638a9cb9))
* update workflow to use ubuntu-latest ([f31657f](https://github.com/nikkbh/social-app/commit/f31657fd6fc3cbb6e0fb8f676de161bb323cdf91))
* user registration to defaul to user role ([a207248](https://github.com/nikkbh/social-app/commit/a207248e068e7ecd2d93ab95b1e7270e11c94d15))
* yaml formatting error ([b12831a](https://github.com/nikkbh/social-app/commit/b12831a3fe70c447960266b0786b795256a3fc43))

## [1.2.1](https://github.com/nikkbh/social-app/compare/v1.2.0...v1.2.1) (2025-09-06)


### Bug Fixes

* deploy yaml file ([3bb19cd](https://github.com/nikkbh/social-app/commit/3bb19cd98d74f4c47d90b1d0d4bb77dcda62b247))
* get version job in yaml files ([f161d9b](https://github.com/nikkbh/social-app/commit/f161d9ba5d8150950b224a8c14849c6ab11cb27a))

## [1.2.0](https://github.com/nikkbh/social-app/compare/v1.1.0...v1.2.0) (2025-09-06)


### Features

* add terraform for IAC and deploy pipelines ([dbb6fad](https://github.com/nikkbh/social-app/commit/dbb6fadacce62f28d3fe2422a92a2b99a8c938a0))

## [1.1.0](https://github.com/nikkbh/social-app/compare/v1.0.0...v1.1.0) (2025-09-02)


### Features

* add workflow to automatically update version and release on push ([c2f69d5](https://github.com/nikkbh/social-app/commit/c2f69d5e37b0b5d94b6b32af9e759f9da02ee80a))

## 1.0.0 (2025-09-02)


### Features

* add automation workflow ([c6ae239](https://github.com/nikkbh/social-app/commit/c6ae2399aec94a5c06059f01ce20c18840336cf0))
* add basic auth for health endpoint ([5406f84](https://github.com/nikkbh/social-app/commit/5406f84674e7e586457df51b4997522bb89a5747))
* add comment creation endpoint to post route ([cc0b340](https://github.com/nikkbh/social-app/commit/cc0b3408b2ac1fae63d3bc37e132538804b2a288))
* add CORS middleware ([d1ee0a3](https://github.com/nikkbh/social-app/commit/d1ee0a31a3e9a2a3fdc6a62d630fdc652e5ea074))
* add create and get post by id ([54096fa](https://github.com/nikkbh/social-app/commit/54096fa51211ecc8342be20bd4aa0bcffd1f09b2))
* add delete and update post ([7c6451d](https://github.com/nikkbh/social-app/commit/7c6451d2dee4919e61e24ab4017d7ae60efa36b5))
* add graceful shutdown ([906f286](https://github.com/nikkbh/social-app/commit/906f286de5e395d90a07fed5815f80976c90d206))
* add justfile for simpler development, add comments to posts and fetch comments on getPostsHandler ([052bbec](https://github.com/nikkbh/social-app/commit/052bbec6e7f1a59dbf15acf1e260ab1e710e44c0))
* add pagination and sorting to get user feed ([672e897](https://github.com/nikkbh/social-app/commit/672e89794f2c2080cc802c4b74b4d6fe6f1911ba))
* add rate limiter to application initialization in test utils ([306ae6f](https://github.com/nikkbh/social-app/commit/306ae6f71a9c90b167c9cf10c68432cbdb933f89))
* add ratelimiting ([a0efb4f](https://github.com/nikkbh/social-app/commit/a0efb4f440a8b46dc4c951ad9fd79657aa28a076))
* add redis cache to improve get user by id performance ([818c037](https://github.com/nikkbh/social-app/commit/818c0375c3eb348288eb972b30bd3fbfb7785a40))
* add release-please workflow for automated releases ([81eb10b](https://github.com/nikkbh/social-app/commit/81eb10b632a769f78383ba9db5259408bbdffe2a))
* add role authorization ([bd396e4](https://github.com/nikkbh/social-app/commit/bd396e4da657ca125cea8a5f752073afc31a0df5))
* add search filters on get user feed ([ba008d6](https://github.com/nikkbh/social-app/commit/ba008d638d619fe9720522ead661f32f20c9b611))
* add server metrics ([eae8cc5](https://github.com/nikkbh/social-app/commit/eae8cc59ed21f328be16c03c90a56d3d56dbfd18))
* add storeage, db code with postgress docker-compose file ([3de8111](https://github.com/nikkbh/social-app/commit/3de811198dfe55171a37f3319ac5eeb8340ab50f))
* add swagger docs and strucutred logging ([589b596](https://github.com/nikkbh/social-app/commit/589b5963083b0b9706ee0244c40c806a2e2985fb))
* add testing setup and test case for get user handler ([575372c](https://github.com/nikkbh/social-app/commit/575372c81bb5d269bf172583f950d4cd782c79a1))
* add user feed endpoint, add indexing to db ([767d0d2](https://github.com/nikkbh/social-app/commit/767d0d2823940d7a4001efe83971894745e2010d))
* add user follow and unfollow endpoint ([76ba579](https://github.com/nikkbh/social-app/commit/76ba579a7d3480046954d9290df17b40e833e1b1))
* added JWT authentication to all protected resources ([3afc4fb](https://github.com/nikkbh/social-app/commit/3afc4fb7970d3e8707877b19398181a4e00e8b60))
* added user activation endpoint via token ([2639682](https://github.com/nikkbh/social-app/commit/263968201055a94f4aa07c09d113e431d498e919))
* base http server setup with hot reloading ([666f4d2](https://github.com/nikkbh/social-app/commit/666f4d298991b4c26a37d3039b372ac89766d7a3))
* create user invitation data layer logic and tables ([8d04b29](https://github.com/nikkbh/social-app/commit/8d04b2904fb9283969685fc6d010b8ccde3c1e56))
* handle race condition, add seed db script, add comment to post ([4402770](https://github.com/nikkbh/social-app/commit/44027704179ea08274469ca6981bdba0624e159f))
* integrated mailtrap to send user registration emails ([fb39220](https://github.com/nikkbh/social-app/commit/fb39220389c2d3bfec035f37f12bf2b89cf5efd7))


### Bug Fixes

* swagger doc to access postID from URLParams ([a18aead](https://github.com/nikkbh/social-app/commit/a18aeada24a25cedd8a136702aac66e5638a9cb9))
* update workflow to use ubuntu-latest ([f31657f](https://github.com/nikkbh/social-app/commit/f31657fd6fc3cbb6e0fb8f676de161bb323cdf91))
* user registration to defaul to user role ([a207248](https://github.com/nikkbh/social-app/commit/a207248e068e7ecd2d93ab95b1e7270e11c94d15))
* yaml formatting error ([b12831a](https://github.com/nikkbh/social-app/commit/b12831a3fe70c447960266b0786b795256a3fc43))

# Installing go dep ?
curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# How to consume API ?
Wikux.postman_collection.json

# Config ?
cp config/config.go config/config.local.go

# Run env ?
docker-compose up

# Run project ?
cd /go/src/wikux-api && go run .

# URL ?
[GIN-debug] GET    /                         --> main.main.func1 (3 handlers)
[GIN-debug] POST   /User/LoginCheck          --> wikux-api/handler.UserLoginCheck (3 handlers)
[GIN-debug] POST   /User/ForgotPassword      --> wikux-api/handler.UserForgotPassword (3 handlers)
[GIN-debug] POST   /User/Registration        --> wikux-api/handler.UserRegistration (3 handlers)
[GIN-debug] GET    /Wikufest/CourseSessionList --> wikux-api/handler.Root (3 handlers)
[GIN-debug] GET    /Wikufest/CourseSessionEnrollmentList --> wikux-api/handler.Root (3 handlers)
[GIN-debug] POST   /Wikufest/CourseSessionEnrollment --> wikux-api/handler.Root (3 handlers)
[GIN-debug] GET    /Wikusamacup/SportList    --> wikux-api/handler.WikusamacupSportList (3 handlers)
[GIN-debug] GET    /Wikusamacup/SportTeamMatchScoreList --> wikux-api/handler.WikusamacupSportTeamMatchScoreList (3 handlers)
[GIN-debug] POST   /Wikusamacup/SportTeamCreate --> wikux-api/handler.WikusamacupSportTeamCreate (3 handlers)
[GIN-debug] POST   /Wikusamacup/SportTeamMatchCreate --> wikux-api/handler.WikusamacupSportTeamMatchCreate (3 handlers)
[GIN-debug] POST   /Wikusamacup/SportTeamMemberCreate --> wikux-api/handler.WikusamacupSportTeamMemberCreate (3 handlers)
[GIN-debug] POST   /Wikusamacup/SportTeamMemberScoreCreate --> wikux-api/handler.WikusamacupSportTeamMemberScoreCreate (3 handlers)
[GIN-debug] GET    /Wikusama/WordpressPostList --> wikux-api/handler.WikusamaWordpressPostList (3 handlers)
[GIN-debug] POST   /Service/SendEmail        --> wikux-api/handler.ServiceSendEmail (3 handlers)

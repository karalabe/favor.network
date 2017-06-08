.PHONY: start-emulators deploy-emulators

start-emulators:
	@echo "Starting Android emulators..."
	$(ANDROID_HOME)/tools/emulator -avd Peter  > /dev/null 2> /dev/null &
	$(ANDROID_HOME)/tools/emulator -avd Jarrad > /dev/null 2> /dev/null &
	$(ANDROID_HOME)/tools/emulator -avd Nick   > /dev/null 2> /dev/null &

deploy-emulators:
	(cd statusbot && http-server &)

	@echo "Deploying chat-bot to Status emulators..."
	$(ANDROID_HOME)/platform-tools/adb -s emulator-5554 forward tcp:5561 tcp:5561
	$(ANDROID_HOME)/platform-tools/adb -s emulator-5554 reverse tcp:8080 tcp:8080
	status-dev-cli remove favor --ip localhost
	status-dev-cli add '{"whisper-identity": "favor",  "name": "Favor Network" ,"bot-url": "http://127.0.0.1:8080/bot.js", "photo-path":"https://www.dropbox.com/s/vzam3sb3nzbs505/robot_scott.png?raw=1"}' --ip localhost

	$(ANDROID_HOME)/platform-tools/adb -s emulator-5556 forward tcp:5561 tcp:5561
	$(ANDROID_HOME)/platform-tools/adb -s emulator-5556 reverse tcp:8080 tcp:8080
	status-dev-cli remove favor --ip localhost
	status-dev-cli add '{"whisper-identity": "favor",  "name": "Favor Network" ,"bot-url": "http://127.0.0.1:8080/bot.js", "photo-path":"https://www.dropbox.com/s/vzam3sb3nzbs505/robot_scott.png?raw=1"}' --ip localhost

	$(ANDROID_HOME)/platform-tools/adb -s emulator-5558 forward tcp:5561 tcp:5561
	$(ANDROID_HOME)/platform-tools/adb -s emulator-5558 reverse tcp:8080 tcp:8080
	status-dev-cli remove favor --ip localhost
	status-dev-cli add '{"whisper-identity": "favor",  "name": "Favor Network" ,"bot-url": "http://127.0.0.1:8080/bot.js", "photo-path":"https://www.dropbox.com/s/vzam3sb3nzbs505/robot_scott.png?raw=1"}' --ip localhost

	killall node

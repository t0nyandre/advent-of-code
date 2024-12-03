.PHONY: new
-include .env
YEAR = $(shell env TZ=Europe/Oslo date +"%Y")
DAY = $(shell env TZ=Europe/Oslo date +"%-d")
LANG = "golang"

new:
	@mkdir -p $(YEAR)/$(LANG)/$(DAY)
	@cp template/$(LANG)/day.tmpl $(YEAR)/$(LANG)/$(DAY)/main.go
	@cp template/$(LANG)/day_test.tmpl $(YEAR)/$(LANG)/$(DAY)/main_test.go
	@touch $(YEAR)/$(LANG)/$(DAY)/example.txt
	@sed -i 's/Day0/Day$(DAY)/g' $(YEAR)/$(LANG)/$(DAY)/main_test.go
	@sed -i 's/Day0/Day$(DAY)/g' $(YEAR)/$(LANG)/$(DAY)/main.go
	@sed -i 's/AOC2024/AOC$(YEAR)/g' $(YEAR)/$(LANG)/$(DAY)/main.go
	@echo "Created AOC$(YEAR) Day$(DAY) from ($(LANG)) template."
	@make scrape-input

scrape-input:
	@curl -s https://adventofcode.com/$(YEAR)/day/$(DAY)/input \
		-b "session=${AOC_SESSION}" \
		> $(YEAR)/$(LANG)/$(DAY)/input.txt
	@echo "Successfully scraped input for $(YEAR)/$(DAY)!"
.PHONY: new
-include .env
YEAR = $(shell env TZ=Europe/Oslo date +"%Y")
DAY = $(shell env TZ=Europe/Oslo date +"%-d")

new:
	@mkdir -p $(YEAR)/$(DAY)
	@cp template/day.tmpl $(YEAR)/$(DAY)/main.go
	@cp template/day_test.tmpl $(YEAR)/$(DAY)/main_test.go
	@touch $(YEAR)/$(DAY)/example.txt
	@sed -i 's/Day0/Day$(DAY)/g' $(YEAR)/$(DAY)/main_test.go
	@sed -i 's/Day0/Day$(DAY)/g' $(YEAR)/$(DAY)/main.go
	@sed -i 's/AOC2024/AOC$(YEAR)/g' $(YEAR)/$(DAY)/main.go
	@echo "Created AOC$(YEAR) Day$(DAY) from template."
	@make scrape-input

scrape-input:
	@curl -s https://adventofcode.com/$(YEAR)/day/$(DAY)/input \
		-b "session=${AOC_SESSION}" \
		> $(YEAR)/$(DAY)/input.txt
	@echo "Successfully scraped input for $(YEAR)/$(DAY)!"
let wFactorList = [0];
function addWhiteFactor() {
	const wFactorElement = document.getElementById("white_factor_items");
	const whiteFactor = document.getElementById("white_factor");
	wFactorList.push(whiteFactor.value);
	wFactorElement.value = wFactorList;
	document.getElementById("white_factor_list").textContent += whiteFactor.selectedOptions[0].label + '\n';
	whiteFactor.value = "";
}

let raceList = [];
function addRaceResults() {
	const race = document.getElementById("race");
	const raceDetail = race.selectedOptions[0].label;
	const raceId = Number(race.value);
	const weather = document.getElementById("race_weather").value;
	const condition = document.getElementById("race_condition").value;
	const tactics = document.getElementById("race_tactics").value;
	const result = Number(document.getElementById("result").value);
	raceList.push({"race_id":raceId, "weather":weather, "condition":condition, "tactics":tactics, "result":result});
	document.getElementById("race_results").value = JSON.stringify(raceList);
	document.getElementById("race_result_list").textContent += raceDetail + ' ' + weather + ' ' + condition + ' ' + tactics + ' ' + result + '着\n';
}

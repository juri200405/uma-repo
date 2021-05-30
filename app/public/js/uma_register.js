let wFactorList = [];
let raceList = [];

window.onload = function() {
	const we = document.getElementById("white_factor_items");
	wFactorList = JSON.parse(we.value).slice(1);
	we.value = JSON.stringify(wFactorList);

	const re = document.getElementById("race_results");
	raceList = JSON.parse(re.value).slice(1);
	re.value = JSON.stringify(raceList);
}

function addWhiteFactor() {
	const whiteFactor = document.getElementById("white_factor");

	wFactorList.push(Number(whiteFactor.value));
	document.getElementById("white_factor_items").value = JSON.stringify(wFactorList);

	const newRow = document.getElementById("white_factor_table").tBodies[0].insertRow();
	newRow.insertCell().innerHTML = `<input type="radio" name="wFactor" value="${whiteFactor.value}">`;
	newRow.insertCell().appendChild(document.createTextNode(whiteFactor.selectedOptions[0].label));
}

function removeWhiteFactor() {
	const radio = document.getElementById("uma").wFactor;
	let index = -1;
	if (radio !== undefined) {
		if (radio.length !== undefined) {
			for (let i = 0; i < radio.length; i++) {
				if (radio[i].checked) {
					index = i;
					break;
				}
			}
		} else if (radio.checked) {
			index = 0;
		}
	}

	if (index > -1) {
		document.getElementById("white_factor_table").tBodies[0].deleteRow(index);

		wFactorList.splice(index, 1);
		document.getElementById("white_factor_items").value = JSON.stringify(wFactorList);
	}
}

function addRaceResults() {
	const race = document.getElementById("race");

	const raceDetail = race.selectedOptions[0].label;
	const raceId = Number(race.value);
	const weather = document.getElementById("race_weather").value;
	const condition = document.getElementById("race_condition").value;
	const tactics = document.getElementById("race_tactics").value;
	const result = Number(document.getElementById("result").value);

	raceList.push({"id":0, "race_id":raceId, "weather":weather, "condition":condition, "tactics":tactics, "result":result});
	document.getElementById("race_results").value = JSON.stringify(raceList);

	const newRow = document.getElementById("race_result_table").tBodies[0].insertRow();
	newRow.insertCell().innerHTML = `<input type="radio" name="raceResult" value="${raceId}">`;
	newRow.insertCell().appendChild(document.createTextNode(raceDetail));
	newRow.insertCell().appendChild(document.createTextNode(weather));
	newRow.insertCell().appendChild(document.createTextNode(condition));
	newRow.insertCell().appendChild(document.createTextNode(tactics));
	newRow.insertCell().appendChild(document.createTextNode(result));
}

function removeRaceResult() {
	const radio = document.getElementById("uma").raceResult;
	let index = -1;
	if (radio !== undefined) {
		if (radio.length !== undefined) {
			for (let i = 0; i < radio.length; i++) {
				if (radio[i].checked) {
					index = i;
					break;
				}
			}
		} else if (radio.checked) {
			index = 0;
		}
	}

	if (index > -1) {
		document.getElementById("race_result_table").tBodies[0].deleteRow(index);

		raceList.splice(index, 1);
		document.getElementById("race_results").value = JSON.stringify(raceList);
	}
}

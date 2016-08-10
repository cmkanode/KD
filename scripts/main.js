window.onload = function(){
    var oSubmit = document.getElementById("submit");
    oSubmit.addEventListener("click", searchDB);
}

function searchDB(){
    var searchterm = document.getElementById("searchterm").value;
    var url = "http://localhost:8080/title/search/" + searchterm;
    //alert(searchterm);
    if(searchterm != ""){
        callService(url, outputResults);
    }
}

function callService(url, callback) {
    'use strict';
	//console.log(url);
    var data_file = url,
        http_request = new XMLHttpRequest(),
        rawdata,
        jsonObj;
    try {
        // Opera 8.0+, Firefox, Chrome, Safari
        http_request = new XMLHttpRequest();
    } catch (e) {
        // Internet Explorer Browsers
        try {
            http_request = new ActiveXObject("Msxml2.XMLHTTP");
        } catch (e2) {
            try {
                http_request = new ActiveXObject("Microsoft.XMLHTTP");
            } catch (e3) {
                // Something went wrong
                alert("Your browser broke!");
                return false;
            }
        }
    }

    http_request.onreadystatechange  = function () {
        if (http_request.readyState === 4) {
            // Javascript function JSON.parse to parse JSON data
            rawdata = http_request.responseText;
            try{
                //console.log(rawdata);
                jsonObj = JSON.parse(rawdata);
                callback(jsonObj);
            }catch(ex){
                console.log("Unexpected Error: " + ex);
            }
        }
    };
    http_request.open("GET", data_file, true);
    http_request.withCredentials = true;
    http_request.send();
}

function removeChildren(parent){
	try{
		while(parent.childNodes.length > 0){
			parent.removeChild(parent.childNodes[0]);
		}
	}catch(ex){
		//console.log(ex);
	}
}

function outputResults(titles){
    var parent = document.getElementById("results")
    removeChildren(parent);
    // TitleName, ReleaseYear
    //console.log("Number of Returned Titles: " + titles.length);
    if(titles.length > 0){
        //console.log("Create List of Returned Titles");
        var oUl = document.createElement("ul");
        for(i = 0; i < titles.length; i++){
            //console.log("Create List Item");
            var oLi = document.createElement("li");
            var text = titles[i].TitleName + ", " + titles[i].ReleaseYear;
            console.log(text);
            oLi.innerHTML = text;
            oUl.appendChild(oLi);
        }
        parent.appendChild(oUl);
    }
    else{
        console.log("No return values");
    }
}

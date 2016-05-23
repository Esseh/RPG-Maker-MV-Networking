(function() {
	// ====================================================================================
	// Ajax
	// Makes an AJAX request.
	// Take in a JSON object of the following form. Omitting a field results in the default.
	// 	{
	//		method:	"GET", or "POST",	default-"GET"
	//		url:	"Valid URL" 		default "/"
	//		async:	true or false		default true
	//		send:	"Sent Info"			default ""
	//	}
	// Another object is taken in for a response. (Optional)
	//	{
	//		ret: Object Variable	If exists it will bind the response text to the variable under the field 'ret'
	//		callback: function		If exists it will use the callback function.
	//		parameters: [...]		Must exist alongside a callback. It will use the callback but applying each of the parameters in order.
	//		use_response: true		Includes the response text as first parameter of the callback. 
	//  }
	// ===================================================================================
	ajax = function(req,res){
		if(res 			 === undefined){ res = {}; }				// If no response specified then make an empty object. 
		if(req["url"]    === undefined){ req["method"] = "/"; }		// If no URL specified default to relative location.
		if(req["method"] === undefined){ req["method"] = "GET"; }	// If no method set assume GET.
		if(req["send"]   === undefined){ req["send"] = ""; }		// If no send set assume empty.
		if(req["async"]  === undefined){ req["async"] = true; }		// If not specified then assume asynchronous.
		var xhttp = new XMLHttpRequest(); // Open XML HTTP Request
		if(res != '{}'){ // Is a response set?
			xhttp.onreadystatechange = function() {	// When ready...
			if (xhttp.readyState == 4 && xhttp.status == 200) {
				// If it exists tie the response text into an object.
				if(res["ret"]        !== undefined){ res["ret"].ret = xhttp.responseText; }	
				// If parameters weren't specified default to an empty array.
				if(res["parameters"] === undefined){ res["parameters"] = []; }
				if(res["use_response" !== undefined]){
					res["parameters"] = [xhttp.responseText].concat(res["parameters"])
				}
				// If a callback was specified then execute it with optional parameters.
				if(res["callback"]   !== undefined){ res["callback"].apply({},res["parameters"]); } 
			}		
		}
	  }
	  // Make the request.
	  xhttp.open(req["method"], req["url"], req["async"]);
	  xhttp.send(req["send"]);
	};
	
	function SMN(){};
	$smn = new SMN();



})();
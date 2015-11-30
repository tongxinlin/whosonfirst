var sock=null; 
var wsuri ="ws://localhost:8080/chat";

window.onload = function(){
	sock = new WebSocket(wsuri);

	sock.onmessage=function(e){
		document.getElementById("list").innerHTML += e.data;
		var div = document.getElementById("list")
		div.scrollTop = div.scrollHeight;
	}
}
function send () {
	var msg = document.getElementById('msg').value;
	if (msg.length == 0){
		document.getElementById('msg').focus();
		return
	}
	var data = msg + "<br/>";
	document.getElementById('msg').value="";
	var div = document.getElementById("list")
	var message = "Costello: " + data
	div.innerHTML += message;
	div.scrollTop = div.scrollHeight; 
	sock.send(data);
}
const ws = new WebSocket("localhost:8050/auction/wsconn");

ws.onopen() = function(event){
    console.log("connected");
}

ws.onerror = function(err){
    console.log(err);
    alert("Something went wrong");
}

var suggestionsTableBody = document.getElementById("suggestions-table");

ws.onmessage() = function(msg){
    var message = JSON.parse(msg.data);

    var newSuggestionRow = suggestionsTableBody.insertRow();
    var priorityCell = newSuggestionRow.insertCell();
    var usernameCell = newSuggestionRow.insertCell();
    var suggestionCell = newSuggestionRow.insertCell();

    let priority = document.createTextNode(message.priority);
    priorityCell.appendChild(priority);

    let username = document.createTextNode(message.username);
    usernameCell.appendChild(username);

    let suggestion = document.createTextNode(message.suggestion);
    suggestionCell.appendChild(suggestion);                                                                                 
}

var suggestion_field = document.getElementById("suggestion-price");

function sendSuggestion(){
    var suggestion = {price : suggestion_field.value};

    ws.send(JSON.stringify(suggestion));
}
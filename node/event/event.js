class EventProxy {
    emit(selected, results) {

    }
}

let proxy = new EventProxy();
let status = "ready";
let select = function(callback){
    proxy.once("selected",callback);
    if(status === "ready"){
        status = "pending";
        var db;
        db.select("SQL", function(results){
            proxy.emit("selected",results);
            status = "ready";
        });
    }};
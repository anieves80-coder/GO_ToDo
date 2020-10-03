$(document).ready(function () {

    const dt = new Date();
    const date = dt.getMonth() + 1 + "/" + dt.getDate() + "/" + dt.getFullYear();
    $("#inputDate").val(date)

    $("#frm").on("submit", function(e){
        e.preventDefault();

        const data = {            
            date: $("#inputDate").val(),
            info: $("#inputBox").val()
        }
    
        jData = JSON.stringify(data);
        
        $.post("/addRec",jData, function(res){
            res = JSON.parse(res);
            if(res.msg === "ok"){                      
                location.reload();
            }
        })

    });    
    
});
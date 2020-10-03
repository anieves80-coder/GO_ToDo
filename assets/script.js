$(document).ready(function () {

    const dt = new Date();
    const date = dt.getMonth() + 1 + "/" + dt.getDate() + "/" + dt.getFullYear();
    
    $("#addBtn").on("click", function(){    

        const data = {
            id: "123456ok",
            date,
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
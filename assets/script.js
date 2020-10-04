$(document).ready(function () {

    let updateId;
    const dt = new Date();
    const date = dt.getMonth() + 1 + "/" + dt.getDate() + "/" + dt.getFullYear();
    $("#inputDate").val(date)

    $("#frm").on("submit", function (e) {
        e.preventDefault();
        if ($("#addBtn").text() === "Update")
            updateRec(updateId)
        else
            addRec()
    });

    $(".btn-del").on("click", function () {
        jData = JSON.stringify({ id: $(this).attr('data-passed_id') });
        $.post("/delRec", jData, function (res) {
            res = JSON.parse(res);
            if (res.msg === "ok") {
                location.reload();
            }
        })
    });

    $(".btn-edit").on("click", function () {

        updateId = $(this).attr('data-passed_id');
        $("#inputBox").val($(this).attr('data-passed_desc'));
        $("#inputDate").val($(this).attr('data-passed_d'));
        $("#btnCancel").removeAttr('hidden');
        $("#addBtn").html("Update");

    });

    $("#btnCancel").on("click", function () {
        location.reload();
    });

});

function addRec() {

    const data = {
        date: $("#inputDate").val(),
        info: $("#inputBox").val()
    }

    jData = JSON.stringify(data);

    $.post("/addRec", jData, function (res) {
        res = JSON.parse(res);
        if (res.msg === "ok") {
            location.reload();
        }
    });

}
function updateRec(id) {

    const data = {
        id,
        date: $("#inputDate").val(),
        info: $("#inputBox").val()
    }

    jData = JSON.stringify(data);

    $.post("/updateRec", jData, function (res) {
        res = JSON.parse(res);
        if (res.msg === "ok") {            
            location.reload();
        }
    })

}
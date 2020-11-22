$(document).ready(function(){

    $(".mobile_nav").click(function(){
        if( $(".mobile_desc").is(":visible") ) {
            $(".mobile_desc").slideUp();
        } else {
            $(".mobile_desc").slideDown();
        }
    });

    $(".open_footer").click(function() {
        if( $(".footer").is(":visible")) {
            $(".footer").slideUp();
        } else {
            $(".footer").slideDown();
        }
    });

    $(".log_out").click(function() {
        $.ajax({
            url: "/logOut",
            type: "get",
            dataType:"json",
            success: function(data){
                console.log("log_out success!");
            }
        });
    });


    /*
    $(window).scroll(function() {
        if ($(window).scrollTop() == $(document).height() - $(window).height()) { 
            $(".footer").slideDown();
        } 
    });​
    */

    /*
   $.ajax({
    url: "/ajax",
    type: "get",
    dataType:"json", //전송받을 형식 지정
    //async: false,     //동기 비동기 설정
    //data : $("#form1").serialize(),  //폼데이터를 직렬화해서 전송 폼전체를 전송시
    //data : {"UserID":"1","Name":"2","Email":"3"},
    //data : params, //파라미터형태로 전송할 경우
    success: function(data){
        var list = $.parseJSON(data);
        console.log(data);
        console.log("-------------");
        console.log(list);

        var listLen = list.length;
        var contentStr = "";
        for(var i=0; i<listLen; i++){
            contentStr +=  list[i]+ "||";
        }
        console.log("------------");
        console.log(contentStr);
        
    }
    });
    */
});
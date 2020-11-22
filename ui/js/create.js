$(document).ready(function(){

    function removeAnimationBlock() {

        $("#loading_animation_block").css({"display":"none"});
        //$("#menu_wrap").css({"display":"block"});
        //$(".map_wrap").css({"display":"block"});
        //$(".map").css({"display":"block"});
        
        //console.log("refreshTabFunc click."); // HTML Console에 Log 출력. IE에서는 console 에러발생.
    }
    removeAnimationBlockFunc = removeAnimationBlock; // javascript 함수에 jQuery 함수 대입.

    $(".search_box").click(function(){
        
            if( $(".map_wrap").is(":visible") ) {
                $(".map_wrap").slideUp();
            } else {
                $('html').scrollTop(0);
                $(".map_wrap").slideDown('fast', function() {
                    //getLoacation();
                    newLocation();
                });
            }
            
        /*
        $.when( ShowMap() ).done(function() {
            setTimeout(function() {
                reloadLayout();
              }, 1000);
            
        });
        */
    });
    
    

    $(".close_kakaomap").click(function(){
        $(".map_wrap").slideUp();
    })
})
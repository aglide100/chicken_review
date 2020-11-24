$(document).ready(function(){
    $("#author").val($("#CheckUser").val());


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
                KakaoMapClose();
            } else {
                $('html').scrollTop(0);
                $(".map_wrap").slideDown('fast', function() {
                    //getLoacation();
                    
                    $('html, body').addClass('hidden');
                    $('body').addClass('scrollDisable').on('scroll touchmove mousewheel', function(e){
                        e.preventDefault();
                    });

                    newLocation();
                });
            }
           
    });

    function KakaoMapClose() {
        if( $(".map_wrap").is(":visible") ) {
            $(".map_wrap").slideUp();
            $('body').removeClass('scrollDisable').off('scroll touchmove mousewheel');
            $('body').css({"overflow":"scroll"});
        }
    }
    

    $(".close_kakaomap").click(function(){
        KakaoMapClose();
    })
})
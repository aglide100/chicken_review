$(document).ready(function(){
    if ($("#CheckUser").val() == "<empty>") {
        $("#author").val("guest");    
    } else {
        $("#author").val($("#CheckUser").val());
    }
    

    function removeAnimationBlock() {
        $("#loading_animation_block").css({"display":"none"});
        
    }
    removeAnimationBlockFunc = removeAnimationBlock; // javascript 함수에 jQuery 함수 대입.
    KakaoMapCloseFunc = KakaoMapClose;

    $(".search_box").click(function(){
        
            if( $(".map_wrap").is(":visible") ) {
                KakaoMapClose(false);
            } else {
                $('html').scrollTop(0);
                $(".map_wrap").slideDown('fast', function() {
                    $('body').addClass('blockScroll');
                    $('body').addClass('scrollDisable').on('scroll touchmove mousewheel', function(e){
                        e.preventDefault();
                    });

                    newLocation();
                });
            }
           
    });

    function KakaoMapClose(checkPushStore) {
        if( $(".map_wrap").is(":visible") ) {
            $(".map_wrap").slideUp();
            if (!checkPushStore) {
                $("#store_name").val("");
                $("#phone_number").val("");
                $("#addr").val("");
                $("#lat").val("<empty>");
                $("#lng").val("<empty>");
            }

            $('body').removeClass('blockScroll');
            $('body').removeClass('scrollDisable').off('scroll touchmove mousewheel');
            $('body').css({"overflow":"scroll"});
        }
    }
    

    $(".close_kakaomap").click(function(){
        KakaoMapClose(false);
    })

    function mapsConfrim() {
        alert("Confirm!");
    }

    function mapsCancel() {

    }
})
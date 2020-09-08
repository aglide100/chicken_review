
//var swiper;


function deleteReview() {
    var check = confirm("delete review?");
    if(check == true) {
        var para = document.location.href.split("delete/");
        var URI = "/delete/";
        URI = URI + para;
        location.href = URI;
    }
}



//function forDesktopSlider() {
    var swiper = new Swiper('.swiper-container', {
        pagination: {
          el: '.swiper-pagination',
          type: 'fraction',
        },
        navigation: {
          nextEl: '.swiper-button-next',
          prevEl: '.swiper-button-prev',
        },
      });

//}

function forMobileSlider() {
    var swiper = new Swiper('.swiper-container', {
        pagination: {
          el: '.swiper-pagination',
          dynamicBullets: true,
        },
      });
}

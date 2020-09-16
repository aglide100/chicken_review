
//var swiper;


function deleteReview() {
    var check = confirm("delete review?");
    if(check == true) {
        var para = document.location.href.split("reviews/");
        console.log(para)
        var URI = "delete/";
        const str = [para[0], URI, para[1]].join('');
        console.log(str)
        location.href = str;
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


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

var swiper = new Swiper('.swiper-container', {
  slidesPerView: 3,
  spaceBetween: 30,
  pagination: {
    el: '.swiper-pagination',
    clickable: true,
  },
});
$(function(){
  var nb = $('.navbtn');
  var n = $('.mainnav');

  $(window).on('resize', function(){

    if($(this).width() < 570 && n.hasClass('keep-nav-closed')) {
      // if the nav menu and nav button are both visible,
      // then the responsive nav transitioned from open to non-responsive, then back again.
      // re-hide the nav menu and remove the hidden class
      $('.mainnav').hide().removeAttr('class');
    }
    if(nb.is(':hidden') && n.is(':hidden') && $(window).width() > 569) {
      // if the navigation menu and nav button are both hidden,
      // then the responsive nav is closed and the window resized larger than 560px.
      // just display the nav menu which will auto-hide at <560px width.
      $('.mainnav').show().addClass('keep-nav-closed');
    }
  });

  // $('.mainnav a,.top a,#btmnav nav a').on('click', function(e){
  //   e.preventDefault(); // stop all hash(#) anchor links from loading
  // });

  $('.navbtn').on('click', function(e){
    e.preventDefault();
    $(".mainnav").toggleClass('open');
  });


  $('.menu').click(function(){
    $(this).toggleClass('open');
    $('.slidenav').toggleClass('open')
    $('body').toggleClass('nav-open')
  })

  // fixed header

  $(window).scroll(function () {
    var sc = $(window).scrollTop()
    if (sc > 1) {
      $(".top").addClass("scroll-on")
    } else {
      $(".top").removeClass("scroll-on")
    }
  });
});

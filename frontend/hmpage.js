let slideIndex = 0;
showSlides();

let slideInterval = setInterval(nextSlide, 3000); // Change the interval (in milliseconds) to adjust the slide speed

function nextSlide() {
  slideIndex++;
  showSlides();
}

function prevSlide() {
  slideIndex--;
  showSlides();
}

function showSlides() {
  let i;
  let slides = document.getElementsByClassName("carousel-item");
  if (slideIndex > slides.length) {
    slideIndex = 1;
  }
  if (slideIndex < 1) {
    slideIndex = slides.length;
  }
  for (i = 0; i < slides.length; i++) {
    slides[i].style.transform = `translateX(-${(slideIndex - 1) * 100}%)`;
  }
}
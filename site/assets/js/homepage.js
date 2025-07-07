//nav menu
function toggleMenu() {
  const nav = document.getElementById('nav-menu');
  nav.style.display = nav.style.display === 'flex' ? 'none' : 'flex';
}


// dot navigation
window.addEventListener('load', () => {
  const scrollWrapper = document.getElementById('services-slider');
  const dots = document.querySelectorAll('#slider-dots .dot');
  const cards = document.querySelectorAll('.service-card');

  function updateDots() {
    const cardWidth = cards[0].offsetWidth;
    const scrollLeft = scrollWrapper.scrollLeft + cardWidth / 2;
    const index = Math.min(cards.length - 1, Math.floor(scrollLeft / cardWidth));
    dots.forEach((dot, i) => {
      dot.classList.toggle('active', i === index);
    });
  }

  scrollWrapper.addEventListener('scroll', () => {
    requestAnimationFrame(updateDots);
  });

  dots.forEach((dot, i) => {
    dot.addEventListener('click', () => {
      const cardWidth = cards[0].offsetWidth;
      scrollWrapper.scrollTo({
        left: i * cardWidth,
        behavior: 'smooth'
      });
    });
  });

  // Scroll to card 2 (index 1) on load
  scrollWrapper.scrollTo({
    left: cards[1].offsetLeft,
    behavior: 'smooth'
  });
});
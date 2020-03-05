function getSlide() {
  return parseInt(window.location.hash.slice(1)) || 0
}

function setSlide(n) {
  window.location = `#${n}`
}

function slideDiv() {
  const div = document.createElement("div")
  div.hidden = true
  return div
}

function slideNumberDiv(i) {
  const div = document.createElement("div")
  div.className = "slide-number"
  div.textContent = i
  return div
}

const slides = [slideDiv()]

function render() {
  const visibleSlide = getSlide()
  slides.forEach((slide, i) => {
    slide.hidden = i !== visibleSlide
  })
}

const root = document.querySelector("#slides")

while (root.children.length > 0) {
  const e = root.removeChild(root.firstChild)
  if (e.tagName === "HR") {
    slides.push(slideDiv())
  } else {
    slides[slides.length - 1].appendChild(e)
  }
}

slides.forEach((slide, i) => {
  if (i > 0) {
    slide.appendChild(slideNumberDiv(i))
  }
  root.appendChild(slide)
})

render()

document.addEventListener("keydown", e => {
  switch (e.key) {
    case " ":
    case "ArrowRight":
    case "ArrowDown":
      e.preventDefault()
      setSlide(Math.min(getSlide() + 1, slides.length - 1))
      render()
      break
    case "Enter":
    case "ArrowLeft":
    case "ArrowUp":
      e.preventDefault()
      setSlide(Math.max(getSlide() - 1, 0))
      render()
      break
    case "f":
      e.preventDefault()
      const requestFullscreen =
        document.documentElement.requestFullscreen ||
        document.documentElement.webkitRequestFullscreen ||
        document.documentElement.mozRequestFullscreen ||
        document.documentElement.msRequestFullscreen
      requestFullscreen.call(document.documentElement)
      break
  }
})

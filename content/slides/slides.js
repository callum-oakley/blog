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
let visibleSlide = 0
function render() {
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
      visibleSlide = Math.min(visibleSlide + 1, slides.length - 1)
      render()
      break
    case "Enter":
    case "ArrowLeft":
    case "ArrowUp":
      visibleSlide = Math.max(visibleSlide - 1, 0)
      render()
      break
  }
})

let minutes = document.getElementById("minutes"),
  seconds = document.getElementById("seconds"),
  pauseBtn = document.getElementById("pause-btn"),
  startBtn = document.getElementById("start-btn"),
  resetBtn = document.getElementById("reset-btn"),
  increaseBtn = document.getElementById("increase"),
  decreaseBtn = document.getElementById("decrease");
// INFO: Current state
let state = {
  started: false,
  running: false,
  paused: false,
  initialSeconds: 1500,
  currSeconds: 1500,
  timerId: 0,
};
minutes.innerText = parseInt(state.initialSeconds / 60);
seconds.innerText = parseInt(state.initialSeconds % 60);
// TODO: Add this for change timer reactivity
function main() {
  if (!state.started) {
    pauseBtn.setAttribute("disabled", "");
  }
  startBtn.addEventListener("click", startTimer);
  pauseBtn.addEventListener("click", pauseTimer);
  resetBtn.addEventListener("click", resetTimer);
  increaseBtn.addEventListener("click", () => {
    changeSeconds(60);
  });
  decreaseBtn.addEventListener("click", () => {
    changeSeconds(-60);
  });
}
function startTimer() {
  //Change state
  state.started = true;
  state.running = true;
  //Set enabled attributes
  startBtn.setAttribute("disabled", "");
  pauseBtn.removeAttribute("disabled");
  //Start timer
  state.timerId = setInterval(() => {
    state.currSeconds--;
    seconds.innerText = state.currSeconds % 60;
    minutes.innerText = parseInt(state.currSeconds / 60);
  }, 1000);
}
function pauseTimer() {
  state.running = false;
  clearInterval(state.timerId);
  startBtn.innerText = "Resume";
  startBtn.removeAttribute("disabled");
}
function resetTimer() {
  state.started = false;
  state.running = false;
  state.currSeconds = state.initialSeconds;
  clearInterval(state.timerId);
  pauseBtn.setAttribute("disabled", "");
  startBtn.removeAttribute("disabled");
}

function changeSeconds(time) {
  state.currSeconds = state.currSeconds + time;

  minutes.innerText = parseInt(state.currSeconds / 60);
  seconds.innerText = parseInt(state.currSeconds % 60);
}
main();

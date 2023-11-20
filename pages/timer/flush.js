minutes = document.getElementById("minutes");
seconds = document.getElementById("seconds");
pauseBtn = document.getElementById("pause-btn");
startBtn = document.getElementById("start-btn");
resetBtn = document.getElementById("reset-btn");
increaseBtn = document.getElementById("increase");
decreaseBtn = document.getElementById("decrease");
if (state.started || state.running) {
  pauseBtn.setAttribute("disabled", "");
}
main();

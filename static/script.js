// set default theme to dark
const root = document.documentElement;
const saved = localStorage.getItem("theme");

// initial theme: saved choice > OS preference > light
root.dataset.theme = saved || "light";

// toggle theme
const toggleThemeBtn = document.getElementById("toggle-theme");
toggleThemeBtn?.addEventListener("click", () => {
  const root = document.documentElement;
  const moon = toggleThemeBtn.querySelector(".moon");
  const sun = toggleThemeBtn.querySelector(".sun");
  console.log("before", root.dataset.theme);

  root.dataset.theme = root.dataset.theme === "light" ? "dark" : "light";
  moon.classList.toggle("hidden", root.dataset.theme === "dark");
  sun.classList.toggle("hidden", root.dataset.theme === "light");
  console.log("current theme:", root.dataset.theme);
});

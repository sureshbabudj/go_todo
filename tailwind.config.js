module.exports = {
  content: [
    "./cmd/web/**/*.html", // Include HTML files in your Go templates
    "./scripts/**/*.ts", // Include TypeScript files for Tailwind classes
    "./styles/**/*.css", // Include CSS files for Tailwind classes
  ],
  theme: {
    extend: {
      fontFamily: {
        sans: ["Rubik", "sans-serif"], // Add Roboto as the default sans font
      },
    },
  },
  plugins: [
    require("daisyui"), // Include DaisyUI plugin
  ],
};

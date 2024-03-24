<!-- Main container for the iPhone-like layout -->
<link rel="stylesheet" href="https://fonts.googleapis.com/css2?family=Material+Symbols+Outlined:opsz,wght,FILL,GRAD@20..48,100..700,0..1,-50..200" />

<div style="border: 2px solid black; width: 320px; height: 600px; position: relative; background-color: #f0f0f5; margin: 0 auto; box-sizing: border-box;">

  <!-- Top Navigation Bar -->
  <div style="background-color: #000; color: #fff; text-align: center; padding: 10px; box-sizing: border-box; display: flex; justify-content: space-between; align-items: center;">
     <span class="material-symbols-outlined">menu</span> <!-- Burger icon for menu -->
    <div>Title</div> <!-- Middle title -->
    <span class="material-symbols-outlined">search</span> <!-- Search icon -->
  </div>

  <!-- Middle Content (Takes over all available space except for nav bars) -->
  <div style="position: absolute; top: 50px; bottom: 65px; left: 0; right: 0; padding: 10px; text-align: center; overflow-y: auto;">
    <!-- Add your content here -->
    Content goes here
  </div>

  <!-- Bottom Navigation Bar -->
  <div style="position: absolute; bottom: 0; width: 100%; background-color: #000; color: #fff; text-align: center; padding: 15px; box-sizing: border-box; display: flex; justify-content: space-around; align-items: center;">
    <span class="material-symbols-outlined">star</span> <!-- Star icon -->
    <span class="material-symbols-outlined">favorite</span> <!-- Heart icon -->
    <span class="material-symbols-outlined">grade</span> <!-- Grade icon -->
    <span class="material-symbols-outlined">search</span> <!-- Search icon -->
  </div>
</div>

---

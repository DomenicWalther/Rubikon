<script lang="ts">
    export let color: string;
    export let isActiveItem : boolean;

    
    let stroke = isActiveItem ? "none" : color
    let fill = isActiveItem ? color : "none"
</script>

<svg width="80" height="80" viewBox="0 0 80 80" fill={fill} stroke={stroke} xmlns="http://www.w3.org/2000/svg">
<rect width="80" height="80" rx="40" fill="url(#paint0_linear_468_6908)"/>
<circle cx="40.0002" cy="41.818" r="27.6818" fill="white" stroke="black"/>
<ellipse cx="40.4549" cy="49.0909" rx="22.2727" ry="19.0909" fill="#252C62"/>
<path fill-rule="evenodd" clip-rule="evenodd" d="M31.7588 41.4049C31.7716 41.9496 31.7167 42.5047 31.5879 43.0598C30.7774 46.5518 27.3405 48.7506 23.8497 48.053C23.9185 51.0003 25.9656 53.6418 28.97 54.339C32.5296 55.1651 36.0849 52.9493 36.911 49.3897C37.7371 45.8301 35.5212 42.2748 31.9616 41.4487C31.894 41.433 31.8264 41.4184 31.7588 41.4049Z" fill="white"/>
<path fill-rule="evenodd" clip-rule="evenodd" d="M49.3754 41.4049C49.3627 41.9496 49.4176 42.5047 49.5464 43.0598C50.3569 46.5518 53.7938 48.7506 57.2846 48.053C57.2158 51.0003 55.1686 53.6418 52.1643 54.339C48.6047 55.1651 45.0494 52.9493 44.2233 49.3897C43.3972 45.8301 45.6131 42.2748 49.1726 41.4487C49.2402 41.433 49.3078 41.4184 49.3754 41.4049Z" fill="white"/>
<path d="M43.6739 61.2789C43.6112 60.8414 43.443 60.4256 43.1838 60.0674C42.9247 59.7091 42.5822 59.4193 42.1861 59.2227C41.7899 59.0262 41.3519 58.9288 40.9099 58.9391C40.4678 58.9494 40.0349 59.067 39.6484 59.2817C39.262 59.4965 38.9337 59.8019 38.6917 60.1718C38.4497 60.5417 38.3014 60.9649 38.2595 61.4049" stroke="white" stroke-width="2" stroke-linecap="round"/>
<defs>
<linearGradient id="paint0_linear_468_6908" x1="40" y1="0" x2="40" y2="80" gradientUnits="userSpaceOnUse">
<stop stop-color="#4592FA"/>
<stop offset="1" stop-color="#78B1FF"/>
</linearGradient>
</defs>
</svg>

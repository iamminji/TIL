import * as THREE from 'three';

const scene = new THREE.Scene();

// three 에는 camera 종류가 몇개 있다.
// attributes information
// - first is the field of view
// FOV is the extent of the scene that is seen on the diplay at any given moment. the value is drees.
// - second is the aspect ratio
// - next two are the near and far clpping plane.
const camera = new THREE.PerspectiveCamera(80, window.innerWidth / window.innerHeight, 1, 500);

const renderer = new THREE.WebGLRenderer();
renderer.setSize(window.innerWidth, window.innerHeight);
document.body.appendChild(renderer.domElement);

// to create cube
const geometry = new THREE.BoxGeometry(1, 1, 1);
// color
const material = new THREE.MeshBasicMaterial({ color: 0x00ff00 });
// mesh
const cube = new THREE.Mesh(geometry, material);
scene.add(cube);

camera.position.z = 5;

function animate() {
    requestAnimationFrame(animate);

    // 회전 속도
    // run every frame (nomrally 60 times per second)
    cube.rotation.x += 0.01;
    cube.rotation.y += 0.01;

    renderer.render(scene, camera);
}

animate();
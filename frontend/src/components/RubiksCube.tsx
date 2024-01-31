import React, { useEffect, useRef, useState } from "react";
import * as THREE from "three";
import { OrbitControls } from "three/examples/jsm/controls/OrbitControls";
import { Piece } from "../constants/interfaces";
import { getColoredFaces, getRotationFromTensor } from "../utils/utils";

const RubiksCube: React.FC<{ cubeData: Piece[] | null }> = ({ cubeData }) => {
  const cubeSize: number = window.innerWidth / 2000;
  const sceneRef = useRef<HTMLDivElement>(null);
  const scene = useRef(new THREE.Scene());
  const camera = useRef(new THREE.PerspectiveCamera(100, 1.1, 0.1, 1000));
  const renderer = useRef(new THREE.WebGLRenderer({ alpha: true })).current; // Initialize renderer here
  renderer.setClearColor(0x000000, 0);
  const cubes = useRef<THREE.Mesh[]>([]);
  const controls = useRef<OrbitControls | null>(null);

  useEffect(() => {
    // Clear existing cubes from the scene
    while (cubes.current.length) {
      const cubeToRemove = cubes.current.pop();
      if (cubeToRemove) {
        scene.current.remove(cubeToRemove);
        cubeToRemove.geometry.dispose();
        if (Array.isArray(cubeToRemove.material)) {
          cubeToRemove.material.forEach((material) => material.dispose());
        } else {
          cubeToRemove.material.dispose();
        }
      }
    }

    // Set the size, color and check renderer's DOM if already attached
    renderer.setSize(window.innerWidth / 2, window.innerHeight);
    if (sceneRef.current && !sceneRef.current.contains(renderer.domElement)) {
      sceneRef.current.appendChild(renderer.domElement);
    }

    // If there is data
    if (cubeData) {
      console.log("Cube Data Structure", cubeData);
      // Create and position the smaller cubes (representing the Rubik's Cube)
      cubeData.forEach((piece) => {
        // The creation of a single cube
        const geometry = new THREE.BoxGeometry(cubeSize, cubeSize, cubeSize); // Instead of cubes, create just piece faces, the exact needed.

        // Get colored faces
        const pieceFaces = getColoredFaces(piece);

        // Create each cube
        const cube = new THREE.Mesh(geometry, pieceFaces);
        cube.position.set(
          piece.Tensor[0] * (cubeSize + 0.1),
          piece.Tensor[1] * (cubeSize + 0.1),
          piece.Tensor[2] * (cubeSize + 0.1)
        );

        // Determine rotation
        const rotation = getRotationFromTensor(piece.Tensor, piece.PieceType);
        cube.rotation.x = rotation.x;
        cube.rotation.y = rotation.y;
        cube.rotation.z = rotation.z;

        scene.current.add(cube);
        cubes.current.push(cube);
        // Add wireframe outline to the cube with a thicker line
        const edges = new THREE.EdgesGeometry(geometry);
        const line = new THREE.LineSegments(
          edges,
          new THREE.LineBasicMaterial({ color: 0x000000, linewidth: 5 })
        );
        cube.add(line);
      });

      // Camera
      camera.current.position.z = 8;

      // OrbitControls and config
      controls.current = new OrbitControls(camera.current, renderer.domElement);
      // controls.current.enableZoom = false; // Disable zooming
      controls.current.enableDamping = true; // Smooth rotations
      controls.current.minPolarAngle = 0; // Remove vertical rotation limits
      controls.current.maxPolarAngle = Math.PI;
      // Set the target to the center of the cube
      const cornerX = (3 * cubeSize) / 1.2;
      const cornerY = (3 * cubeSize) / 1.2;
      const cornerZ = (3 * cubeSize) / 1.2;
      controls.current.target.set(cornerX / 2, cornerY / 2, cornerZ / 2);
    }

    // Render Animation loop
    function animate() {
      requestAnimationFrame(animate);

      if (controls.current) {
        controls.current.update(); // Update camera controls
      }
      renderer.render(scene.current, camera.current);
    }
    animate();

    // Clean up
    return () => {
      // Dispose of resources when the component unmounts
      scene.current.remove();
    };
  }, [cubeData]);

  return <div id="rubiks-cube-model" ref={sceneRef}></div>;
};

export default RubiksCube;

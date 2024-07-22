<script lang="ts">
	import '../app.css';
	import { Navbar, NavBrand, NavLi, NavUl, NavHamburger, Heading, A } from 'flowbite-svelte';
	import logo from '../lib/logo.png';
</script>

<Navbar class="drop-shadow">
	<NavBrand href="/home">
		<img src={logo} class="h-18 me-3" alt="C4C Logo" />
		<span class="=self-center whitespace-nowrap font-semibold">
			<Heading customSize="text-2xl" color="text-heading" class="font-heading"
				>Elevate 4 Epilepsy {new Date().getFullYear()}</Heading
			>
		</span>
	</NavBrand>
	<NavHamburger />
	<NavUl nonActiveClass="hover:text-hover">
		<NavLi href="https://www.crossoversforcures.org/">Home</NavLi>
		<NavLi href="https://www.crossoversforcures.org/elevate-4-epilepsy/" class="text-l">E4E</NavLi>
		<NavLi href="https://www.crossoversforcures.org/about" class="text-l">About</NavLi>
		<NavLi href="https://www.crossoversforcures.org/contact" class="text-l">Contact</NavLi>
		<NavLi href="https://www.crossoversforcures.org/donate" class="text-l">Donate</NavLi>
	</NavUl>
</Navbar>
<br />

<slot />

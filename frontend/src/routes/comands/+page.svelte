<script lang="ts">
    import { onMount } from 'svelte';
    

    let firstName: string = '';
    let lastName: string = '';
    let patronymic: string = '';
    let group: string = '';
    let role: string = '';
    let techStack: string[] = [];
    let otherTech: string = '';
    let about: string = '';
    let achievements: string = '';
    let isCommander: boolean = false;
  
    const roles = ['Student', 'Junior', 'Middle', 'Senior'];
    const technologies = ['React', 'Vue', 'Svelte', 'Node.js', 'Python', 'Java'];
  
    async function fetchUserProfile() {
      try {
        const response = await fetch('/api/user-profile');
        if (!response.ok) {
          throw new Error('Error fetching user profile');
        }
  
        const data = await response.json();
        
        firstName = data.firstName || '';
        lastName = data.lastName || '';
        patronymic = data.patronymic || '';
        group = data.group || '';
        role = data.role || '';
        techStack = data.techStack || [];
        otherTech = data.otherTech || '';
        about = data.about || '';
        achievements = data.achievements || '';
        isCommander = data.isCommander || false;
      } catch (error: unknown) {
        if (error instanceof Error) {
          console.error('Failed to fetch user profile:', error.message);
        } else {
          console.error('An unknown error occurred');
        }
      }
    }
  
    onMount(() => {
      fetchUserProfile();
    });
  
    async function submitProfile() {
      const profileData = {
        firstName,
        lastName,
        patronymic,
        group,
        role,
        techStack,
        otherTech,
        about,
        achievements,
        isCommander,
      };
  
      try {
        const response = await fetch('/api/update-profile', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(profileData),
        });
  
        if (!response.ok) {
          throw new Error('Error updating profile');
        }
  
        alert('Profile updated successfully!');
      } catch (error: unknown) {
        if (error instanceof Error) {
          alert('Error updating profile: ' + error.message);
        } else {
          alert('Error updating profile: Unknown error');
        }
      }
    }
  </script>

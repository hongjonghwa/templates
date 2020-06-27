 ls -d */ | cut -f1 -d'/' | xargs -I{} git -C {} pull

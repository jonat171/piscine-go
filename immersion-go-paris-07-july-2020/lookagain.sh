find . -name "*.sh" | sed 's|.*/||' | cut -f1 -d '.'

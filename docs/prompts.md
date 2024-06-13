# A prompt to generate long sentences that replicate conversations.

	`Generate a list of %d sentence translations in %s for a %s level lesson on the topic "%s". 
	Return an object with the following structure:
	{
		"lessonType: "The type of lesson",
		"language": "%s",
		"level": "Summary of the lesson difficulty",
		"topic": "The topic of the lesson",
		"description": "Brief description of the lesson",
		"content": [
			{
				"lessonText": ["sentence1", "sentence2", ...],
				"englishText": ["english1", "english2", ...],
				"lessonSyllables": ["word1_syllable1-word1_syllable2 word2_syllable1-word2_syllable2", ...],
				"phoneticSpellings": ["word1_syllable1-word1_syllable2 word2_syllable1-word2_syllable2", ...]
			}
		]
	}
	Rules:
	1. Ensure each sentence is grammatically correct with appropriate spacing and punctuation and no field has trailing spaces.
	2. Ensure indices match across arrays, i.e., sentence1, english1, and its syllables should all correspond to index 0.
	3. Use hyphens ('-') to separate syllables within words in the 'syllables' array but not before punctuation.
	4. Use spaces to separate words within the 'syllables' array.
	5. For Japanese, split syllables by kana or mora, including long vowels and particles as separate units (e.g., 'すみません、メニューをください' should be 'す-み-ま-せ-ん', 'め-にゅー-を-く-だ-さ-い').
	6. For Japanese, clearly distinguish different words in the 'syllables' array with spaces.
	7. Ensure syllables are accurately split according to the %s language's phonetic rules.
	8. Represent each syllable and word correctly without combining multiple syllables into one or improperly breaking single syllables.
	9. Ensure each English translation is contextually and semantically accurate to the corresponding sentence in %s.
	10. phoneticSpellings should be the sounds of the characters in the %s language (phonetic spellings), not the words from the English translation.

	Example Output for Japanese (Beginner Level):
	{
		"lessonType": "Sentence"
		"language": "Japanese",
		"level": "Beginner",
		"topic": "Ordering Food",
		"description": "This lesson covers basic phrases and vocabulary for ordering food in Japanese.",
		"content": [
			{
				"lessonText": ["お水を一つお願いします。", "美味しいランチを注文したいです。", "デザートメニューは何がありますか？", "ご飯とお味噌汁をセットでお願いします。"],
				"englishText": ["One water, please.", "I would like to order a delicious lunch.", "What desserts do you have on the menu?", "I'll have rice and miso soup as a set, please."],
				"lessonSyllables": ["お-み-ず-を ひと-つ お-ね-が-い-し-ま-す", "おい-し-い らん-ち を ちゅう-もん-し-た-い-です", "で-ざー-と め-にゅー- は なに が あ-り-ま-す-か", "ご-はん と お-み-そ-し-る を セット で お-ね-が-い-し-ま-す"],
				"phoneticSpellings": ["o-mi-zu o hi-to-tsu o-ne-ga-i-shi-ma-su", "o-i-shi-i ran-chi o chuu-mon-shi-ta-i de-su", "de-za-a-to me-nyuu wa na-ni ga a-ri-ma-su ka", "go-han to o-mi-so-shi-ru o set-to de o-ne-ga-i-shi-ma-su"]
			}
		]
	}
	Notes:
	- This prompt is designed to generate a lesson object in the specified format for language learning.
	- Make sure that the syllables are split and represented accurately according to the phonetic rules of the target language.
	- Ensure that the phoneticSpellings provide a close phonetic approximation of the pronunciation of the sentences in the lesson language, reflecting how they sound when spoken in that language.`

## example curl

	curl -X POST \
	-H "Content-Type: application/json" \
	-d '{"query": "{ getLesson(lessonType: \"sentence\", language: \"Japanese\", level: \"Beginner\", topic: \"Ordering food\") { lessonType language level description content { ... on WordOrSentenceLesson { lessonText englishText lessonSyllables phoneticSpellings } } } }"}' \
	http://localhost:8080/graphql


# A prompt to generate single words, including the corresponding sound or meaning.

	`Generate a list of %d word translations in %s for a %s level lesson on the topic "%s".
	Return an object with the following structure:
	{
		"lessonType": "%s",
		"language": "%s",
		"level": "Summary of the lesson difficulty",
		"topic": "The topic of the lesson",
		"description": "Brief description of the lesson",
		"content": [
			{
				"lessonText": ["word1", "word2", ...],
				"englishText": ["english1", "english2", ...],
				"lessonSyllables": ["word1_syllable1-word1_syllable2", "word2_syllable1-word2_syllable2", ...],
				"phoneticSpellings": ["word1_phonetic", "word2_phonetic", ...]
			}
		]
	}
	Rules:
	1. Ensure each word is spelled correctly with appropriate spacing and no field has trailing spaces.
	2. Ensure indices match across arrays, i.e., word1, english1, and its syllables should all correspond to index 0.
	3. Use hyphens ('-') to separate syllables within words in the 'syllables' array.
	4. Ensure syllables are accurately split according to the %s language's phonetic rules.
	5. Represent each syllable and word correctly without combining multiple syllables into one or improperly breaking single syllables.
	6. Ensure each English translation is contextually and semantically accurate to the corresponding word in %s.
	7. Phonetic spellings should reflect the sounds of the characters in the %s language (phonetic spellings), not the words from the English translation.
	
	Example Output for Japanese (Basic Level):
	{
		"lessonType": "Word",
		"language": "Japanese",
		"level": "Beginner",
		"topic": "Food Vocabulary",
		"description": "This lesson covers beginner vocabulary related to food in Japanese.",
		"content": [
			{
				"lessonText": ["水", "料理", "デザート", "ご飯"],
				"englishText": ["water", "cooking", "dessert", "rice"],
				"lessonSyllables": ["み-ず", "りょう-り", "で-ざー-と", "ご-はん"],
				"phoneticSpellings": ["mi-zu", "ryou-ri", "de-za-a-to", "go-han"]
			}
		]
	}
	Notes:
	- This prompt is designed to generate a lesson object in the specified format for language learning.
	- Make sure that the syllables are split and represented accurately according to the phonetic rules of the target language.
	- Ensure that the phonetic spellings provide a close phonetic approximation of the pronunciation of the words in the lesson language, reflecting how they sound when spoken in that language.`

## example curl

	curl -X POST \
		-H "Content-Type: application/json" \
		-d '{"query": "{ getLesson(lessonType: \"word\", language: \"Japanese\", level: \"Expert\", topic: \"Ordering food\") { lessonType language level description content { ... on WordOrSentenceLesson { lessonText englishText lessonSyllables phoneticSpellings } } } }"}' \
		http://localhost:8080/graphql


# A prompt for alphabet

	`Generate a random list of %d alphabet characters in %s for a %s level lesson on the topic "%s".
	Return an object with the following structure:
	{
		"lessonType": "%s",
		"language": "%s",
		"level": "Summary of the lesson difficulty",
		"description": "Brief description of the lesson",
		"content": [
			{
				"alphabetCharacter": ["character1", "character2", ...],
				"phoneme": ["phoneme1", "phoneme2", ...]
			}
		]
	}
	Rules:
	1. Ensure each alphabet character is correctly represented and matches the corresponding phoneme.
	2. Ensure indices match across arrays, i.e., character1 and phoneme1 should correspond to index 0.
	3. Use appropriate notation to represent phonemes clearly and accurately for the %s language.
	4. Ensure phonemes are contextually accurate to their corresponding alphabet character in %s.
	5. Represent each alphabet character and phoneme without combining multiple sounds into one or improperly breaking single sounds.
	6. Ensure that each phoneme provides a close phonetic approximation of the pronunciation of the alphabet character in the lesson language.
	7. Do not take into account the topic. As this is focused on alphabet learning it is not relevant.

	Example Output for Japanese (Beginner Level):
	{
		"lessonType": "alphabet",
		"language": "Japanese",
		"level": "Beginner",
		"description": "This lesson covers the basic Hiragana characters and their phonetic sounds in Japanese.",
		"content": [
			{
				"alphabetCharacter": ["あ", "い", "う", "え", "お", ...], // Include all relevant characters
				"phoneme": ["a", "i", "u", "e", "o", ...] // Include phonetic sounds for each character
			}
		]
	}
	Notes:
	- This prompt is designed to generate a lesson object in the specified format for language learning.
	- Make sure that the phonemes are split and represented accurately according to the phonetic rules of the target language.
	- Ensure that the phonemes provide a close phonetic approximation of the pronunciation of the alphabet characters, reflecting how they sound when spoken in that language.`

## example curl

	curl -X POST \
        -H "Content-Type: application/json" \
        -d '{"query": "{ getLesson(lessonType: \"alphabet\", language: \"Japanese\", level: \"Beginner\" topic: \"\") { lessonType language level description content { ... on AlphabetLesson { alphabetCharacter phoneme } } } }"}' \
        http://localhost:8080/graphql

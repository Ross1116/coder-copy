import os
import sys
import json
from typing import Dict, List, Optional, Union
import logging

# Configure logging
logging.basicConfig(
    level=logging.INFO,
    format='%(asctime)s - %(name)s - %(levelname)s - %(message)s'
)
logger = logging.getLogger(__name__)

class DataProcessor:
    """
    A class to process and transform data from various sources.
    
    This class provides methods to load, validate, transform and export data
    in different formats. It supports JSON, CSV and database connections.
    """
    
    def __init__(self, config_path: str = "config.json"):
        # Initialize with default configuration or load from file
        self.config = self._load_config(config_path)
        self.data = None  # Will hold the data once loaded
        self._processed = False  # Track if data has been processed
        
    def _load_config(self, config_path: str) -> Dict:
        """Load configuration from a JSON file.
        
        Args:
            config_path: Path to the configuration file
            
        Returns:
            Dict containing configuration parameters
        """
        try:
            with open(config_path, 'r') as f:
                return json.load(f)
        except FileNotFoundError:
            logger.warning(f"Config file {config_path} not found. Using defaults.")
            return {
                "input_format": "json",
                "output_format": "json",
                "validate": True
            }
    
    def load_data(self, source: str) -> bool:
        """
        Load data from the specified source.
        
        Args:
            source: Path to data file or connection string
            
        Returns:
            True if data was loaded successfully, False otherwise
        """
        # Reset processed flag when loading new data
        self._processed = False
        
        # Determine the type of data source
        if source.endswith('.json'):
            try:
                with open(source, 'r') as f:
                    self.data = json.load(f)
                logger.info(f"Successfully loaded JSON data from {source}")
                return True
            except Exception as e:
                logger.error(f"Failed to load JSON from {source}: {e}")
                return False
        elif source.endswith('.csv'):
            # TODO: Implement CSV loading
            logger.error("CSV loading not yet implemented")
            return False
        else:
            # Handle other data sources
            logger.error(f"Unsupported data source: {source}")
            return False
    
    def process(self) -> Optional[Dict]:
        """Process the loaded data according to configuration settings."""
        if self.data is None:
            logger.error("No data loaded. Call load_data() first.")
            return None
            
        # Apply transformations based on config
        result = self.data  # Start with the original data
        
        # Example transformation: Filter out records with missing values
        if self.config.get("remove_incomplete", False):
            if isinstance(result, list):
                result = [r for r in result if all(r.values())]
        
        '''
        This is a multi-line comment using triple quotes.
        We could add more complex transformations here:
        - Data normalization
        - Feature extraction
        - Aggregation
        - etc.
        '''
        
        self._processed = True
        logger.info("Data processing completed")
        return result
    
    def export(self, destination: str) -> bool:
        """
        Export processed data to the specified destination.
        
        Args:
            destination: Path where data should be saved
            
        Returns:
            True if export was successful, False otherwise
        """
        if not self._processed:
            logger.warning("Data has not been processed. Running process() first.")
            self.process()
            
        if destination.endswith('.json'):
            try:
                with open(destination, 'w') as f:
                    json.dump(self.data, f, indent=2)
                logger.info(f"Successfully exported data to {destination}")
                return True
            except Exception as e:
                logger.error(f"Failed to export to {destination}: {e}")
                return False
        else:
            # Handle other export formats
            logger.error(f"Unsupported export format: {destination}")
            return False


# Example usage of the class
if __name__ == "__main__":
    # Parse command line arguments
    if len(sys.argv) < 3:
        print("Usage: python data_processor.py <input_file> <output_file>")
        sys.exit(1)
        
    input_file = sys.argv[1]
    output_file = sys.argv[2]
    
    # Create processor and run the pipeline
    processor = DataProcessor()
    if processor.load_data(input_file):
        processor.process()
        if processor.export(output_file):
            print(f"Successfully processed {input_file} and saved to {output_file}")
        else:
            print(f"Failed to export to {output_file}")
    else:
        print(f"Failed to load data from {input_file}")